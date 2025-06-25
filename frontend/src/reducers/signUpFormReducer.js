export const initialState = {
  firstName: "",
  lastName: "",
  gender: "Male",
  country: "",
  email: "",
  password: "",
  date: "",
  errors: {
    firstName: false,
    lastName: false,
    country: false,
    email: "",
    password: false,
    date: false,
  },
};

export function Signupreducer(state, action) {
  switch (action.type) {
    case "SET_FIELD":
      return { ...state, [action.field]: action.value };
    case "SET_ERROR":
      return {
        ...state,
        errors: { ...state.errors, [action.field]: action.value },
      };
    default:
      return state;
  }
}
