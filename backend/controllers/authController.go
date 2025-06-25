package controllers

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"gin_app/config"
	"gin_app/dto"
	"gin_app/helpers"
	"gin_app/models"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	userSignUpRequestDTO := dto.UserSignUpRequestDTO{}
	var existingUser models.User

	err := c.ShouldBindJSON(&userSignUpRequestDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid JSON data",
		})
		return
	}

	err = helpers.ValidateSignupData(&userSignUpRequestDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status" : http.StatusBadRequest,
			"error" : err.Error(),
		})
		return
	}

	result := config.DB.Where("email = ?", userSignUpRequestDTO.Email).First(&existingUser)

	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"status":  http.StatusConflict,
			"message": "Email already registered",
		})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userSignUpRequestDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Couldn't hash the password",
		})
		return
	}

	dob, err := time.Parse("2006-01-02", userSignUpRequestDTO.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Couldn't parse the date",
			})
			return
	}

	code, err := rand.Int(rand.Reader, big.NewInt(1000000))

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" : http.StatusInternalServerError,
			"message" : "Couldnt proceed with the request",
		})
		return
	}

	verificationToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email" : userSignUpRequestDTO.Email,
		"country_code" : userSignUpRequestDTO.CountryCode,
	})

	token,err := verificationToken.SignedString([]byte(os.Getenv("VERIFICATION_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" : http.StatusInternalServerError,
			"message" : "Internal Server Error",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Verification", token, 3600*24*30, "", "", true, true)


	type registrationData struct {
		Code *big.Int
		FirstName string 
		LastName string
		DOB time.Time
		CountryCode string
		Email string
		Gender string
		Password string
	}

	data := &registrationData{
		Code: code ,
		FirstName: userSignUpRequestDTO.FirstName,
		LastName: userSignUpRequestDTO.LastName,
		DOB: dob,
		CountryCode: userSignUpRequestDTO.CountryCode,
		Email: userSignUpRequestDTO.Email,
		Gender: userSignUpRequestDTO.Gender,
		Password: string(hashedPassword),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize user data"})
		return
	}
	ctx := context.Background()

	err = config.Client.Set(ctx,"email:verify:" + data.Email,jsonData,5*time.Minute).Err()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" : http.StatusInternalServerError,
			"message" : "Internal Server Error",
		})
		return
	}

	err = 	helpers.SendMail(data.Email,code)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status" : http.StatusInternalServerError,
			"Message" : "Internal Server Error",
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"status" : http.StatusOK,
	})
}


func VerifyRegistration(c *gin.Context) {

	tokenStr,err := c.Cookie("Verification")

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}

	token,err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("VERIFICATION_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.SetCookie("Verification", "", -1, "", "", true, true)
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.SetCookie("Verification", "", -1, "", "", true, true)
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}
	
    type Code struct {
        Code *big.Int `json:"code"`
    }

    var code Code
    if err := c.ShouldBindJSON(&code); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid JSON data",
        })
        return
    }

    email := claims["email"]
	key := "email:verify:" + email.(string);

    val, err := config.Client.Exists(context.Background(), key).Result()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status" : http.StatusInternalServerError,
            "message" : "Internal Server Error",
        })
        return
    }
    if val == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "Url has expired",
        })
        return
    }


    dataObj, err := config.Client.Get(c.Request.Context(), key).Result()
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Token has expired or is invalid",
        })
        return
    }
    type registrationData struct {
        Code        *big.Int  `json:"code"`
        FirstName   string    `json:"firstname"`
        LastName    string    `json:"lastname"`
        DOB         time.Time `json:"dob"`
        CountryCode string    `json:"countrycode"`
        Email       string    `json:"email"`
        Gender      string    `json:"gender"`
        Password    string    `json:"password"`
    }

    data := &registrationData{}
    if err := json.Unmarshal([]byte(dataObj), data); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Internal Server Error",
        })
        return
    }

    if code.Code == nil || data.Code == nil || code.Code.Cmp(data.Code) != 0 {

        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid Code",
        })
        return
    }

	user := models.User{
		FirstName: data.FirstName,
		LastName: data.LastName,
		Gender: data.Gender,
		DateOfBirth: data.DOB,
		CountryCode: data.CountryCode,
		Email: data.Email,
		Password: data.Password,
	}
    tx := config.DB.Begin()

	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback();
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}
	tokenTiger, err := helpers.GetTigerGraphToken()

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	client := resty.New()

	url := fmt.Sprintf("%s/restpp/query/%s/InsertAUser", os.Getenv("TG_HOST"), os.Getenv("TG_GRAPH"))
	res, err := client.R().SetHeader("Authorization", "Bearer "+tokenTiger).SetQueryParams(map[string]string{
		"id":   strconv.FormatUint(uint64(user.ID), 10),
		"name": data.FirstName + " " + data.LastName,
		"age": strconv.FormatInt(int64(time.Now().Year() - data.DOB.Year()),10),
	}).Get(url)

	if err != nil || res.IsError() {
		tx.Rollback();
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	tx.Commit()
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := jwtToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)

	config.Client.Del(context.Background(), key)
	userSignUpResponseDTO := dto.UserSignUpResponseDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"data":    userSignUpResponseDTO,
	})
}

func VerifyTokenUrl(c * gin.Context) {
	tokenStr,err := c.Cookie("Verification")

	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}

	token,err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("VERIFICATION_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		c.SetCookie("Verification", "", -1, "", "", true, true)
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		c.SetCookie("Verification", "", -1, "", "", true, true)
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"message" : "Page not found",
		})
		return
	}
	email := claims["email"]
	key := "email:verify:" + email.(string);

	val, err := config.Client.Exists(context.Background(), key ).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	if val == 0 {
		c.SetCookie("Verification", "", -1, "", "", true, true)
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": "Page not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"valid":  true,
	})
}


func Login(c *gin.Context) {

	userLoginRequestDTO := dto.UserLoginRequestDTO()

	var existingUser models.User
	err := c.ShouldBindJSON(userLoginRequestDTO) 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid Json Data",
		})
		return
	}

	result := config.DB.Where("email = ?", userLoginRequestDTO.LoginField).Find(&existingUser)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Email Or Password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userLoginRequestDTO.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Email Or Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal Server Error",
		})
		return
	}

	c.SetCookie(
    "Authorization",      
    tokenString,         
    3600*24*30,           
    "/",                  
    "localhost",         
    false,                
    true,                 
	)

	userLoginResponseDTO := dto.UserLoginResponseDTO(existingUser.ID, existingUser.FirstName + " " + existingUser.LastName)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data" : userLoginResponseDTO,
	})

}

func CheckLoginStatus(c *gin.Context) {
	tokenString, _ := c.Cookie("Authorization")

	if tokenString != "" {
		claims, err := helpers.VerifyJWT(tokenString)

		if err != nil {
			c.SetCookie("Authorization", "", -1, "", "", true, true)
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Invalid or expired token",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"data": gin.H{
					"ID": claims["sub"],
				},
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  http.StatusUnauthorized,
		"message": "You are not Logged In",
	})
}

func Logout(c *gin.Context) {

	token, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not logged in",
		})
		return
	}

	claims, err := helpers.VerifyJWT(token)

	if err != nil {
		c.SetCookie("Authorization", "", -1, "", "", true, true)

		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid or expired token",
		})
		return
	}

	ctx := context.Background()
	expTime := claims["exp"].(float64) - float64(time.Now().Unix())

	pipe := config.Client.Pipeline()
	pipe.SetNX(ctx, `auth:blacklist:`+token, "", time.Duration(expTime)*time.Second)
	_, err = pipe.Exec(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Could not log out",
		})
		return
	}
	c.SetCookie("Authorization", "", -1, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successfully Logged Out",
	})
}

func Authorization(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}

	claims, err := helpers.VerifyJWT(tokenString)

	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}

	var userContext dto.UserContext


	config.DB.Raw("SELECT id, first_name || ' ' || last_name as fullname, country_code FROM users WHERE id = ?", claims["sub"]).Scan(&userContext)


	
	if userContext.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "You are not authorized to do that",
		})
		c.Abort()
		return
	}
	c.Set("user", userContext);
	c.Next()

}

func CheckEmailExists(c *gin.Context) {

	email := c.Query("email")
	if email == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": "Invalid email",
        })
        return
    }
    
    res := config.DB.Table("users").Where("email = ?", email).Find(&models.User{})
    if res.RowsAffected > 0 {
        c.JSON(http.StatusOK, gin.H{
            "status":  http.StatusOK,
            "message": "Email already exists",
        })
        return
    }

    c.JSON(http.StatusNotFound, gin.H{
        "status":  http.StatusNotFound,
        "message": "Email doesn't exist",
    })
}
