function FormError({ children, size }) {
  return (
    <span className={`text-red-500 ${size ? size : "text-sm"}`}>
      {children}
    </span>
  );
}

export default FormError;
