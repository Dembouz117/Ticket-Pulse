type ValidationResult = string | undefined;

export function validateEmail(email: string): ValidationResult {
    const re = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    if (!re.test(String(email).toLowerCase())) {
        return "Invalid email";
    }
}

export function validatePassword(password: string): ValidationResult {
    const re = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$/;
    if (!re.test(password)) {
        return "Password should contain at least 1 uppercase letter, 1 lowercase letter, and 1 number";
    }
}

// export function validatePhoneNumber(phone: string): ValidationResult {
//     const re = /^[0-9]{8}$/;  // adjust based on your requirements
//     if (!re.test(phone)) {
//         return "Invalid phone number";
//     }
// }