import { object, string } from 'yup'

export const registerValidationSchema = object().shape({
    email: string().trim().required("El correo electrónico es requerido"),
    password: string().trim().required("La contraseña es requerida")
});