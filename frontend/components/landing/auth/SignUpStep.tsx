import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Box, Button, Input, useToast, Text, VStack } from "@chakra-ui/react";
import "react-phone-number-input/style.css";
import PhoneInput, { isValidPhoneNumber } from "react-phone-number-input";
import axios from "axios";
import { getAuthApiUrl } from "@/utilities/common";

interface SignUpFormData {
    email: string;
    password: string;
    confirmPassword: string;
    name: string;
    phoneNumber: string;
}

interface SignUpStepProps {
    onClose: () => void;
    setStep: (step: string) => void;
    setEmail: (email: string) => void;
    setPassword: (password: string) => void;
}

const SignUpStep = ({ onClose ,setStep, setEmail, setPassword }: SignUpStepProps) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<SignUpFormData>();
    const toast = useToast();

    const [phoneNumber, setPhoneNumber] = useState<string>("");
    const [passwordMismatch, setPasswordMismatch] = useState(false);

    const onSubmit = async (data: SignUpFormData) => {
        if (data.password !== data.confirmPassword) {
            setPasswordMismatch(true);
            return;
        } else {
            setPasswordMismatch(false);
        }

        if (isValidPhoneNumber(phoneNumber)) {
            const payload = {
                email: data.email,
                password: data.password,
                name: data.name,
                phone: phoneNumber
            };
            try {
                console.log("Calling API with:", payload);
            const createUserRes = await axios.post(`${getAuthApiUrl()}/user`, payload);
            if (createUserRes.status === 200) {
                const loginRes = await axios.post(`${getAuthApiUrl()}/login`, {
                    email: data.email,
                    password: data.password,
                });
                if (loginRes.status === 200) {
                    setEmail(data.email);
                    setPassword(data.password);
                    setStep("twoFactor");
                }
            }
                
            } catch (error: any) {
                console.log(error.message)
                toast({
                    title: "Error creating acccount, please try again later.",
                    status: "error",
                    duration: 3000,
                    isClosable: true,
                });
                onClose();
            }
            
        } else {
            toast({
                title: "Phone number invalid",
                status: "error",
                duration: 3000,
                isClosable: true,
            });
        }
    };

    return (
        <Box w="full" p={4}>
            <div className="flex flex-col align-middle justify-start gap-4 py-6">
                <div className="text-3xl mb-2">Welcome Back!</div>

                {/* email input */}
                <div className="">
                    <div className="font-light mb-1 text-[#AFAFAF]">Email</div>
                    <Input
                        {...register("email", {
                            required: "Required",
                            pattern: {
                                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
                                message: "invalid email address",
                            },
                        })}
                        placeholder="johndoe@example.com"
                        _placeholder={{ color: "gray" }}
                        border={
                            errors.email ? "1px solid red" : ".5px solid gray"
                        }
                    />
                    {errors.email && (
                        <Text color="red.500">{errors.email.message}</Text>
                    )}
                </div>

                {/* name input */}
                <div>
                    <div className="font-light mb-1 text-[#AFAFAF]">Name</div>
                    <Input
                        {...register("name", { required: "Required" })}
                        placeholder="Name"
                        _placeholder={{ color: "gray" }}
                        border={
                            errors.name ? "1px solid red" : ".5px solid gray"
                        }
                    />

                    {errors.name && (
                        <Text color="red.500">{errors.name.message}</Text>
                    )}
                </div>

                {/* password input */}
                <div>
                    <div className="font-light mb-1 text-[#AFAFAF]">Password</div>
                    <Input
                        {...register("password", {
                            required: "Required",
                            pattern: {
                                value: /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.{8,})/,
                                message:
                                    "Password must contain: 1 uppercase, 1 lowercase, 1 number and be at least 8 characters long",
                            },
                        })}
                        type="password"
                        placeholder="Password"
                        _placeholder={{ color: "gray" }}
                        border={
                            errors.password || passwordMismatch
                                ? "1px solid red"
                                : ".5px solid gray"
                        }
                    />
                    {errors.password && (
                        <Text color="red.500">{errors.password.message}</Text>
                    )}
                </div>

                {/* confirm password input */}
                <div>
                    <div className="font-light mb-1 text-[#AFAFAF]">Confirm Password</div>
                    <Input
                        {...register("confirmPassword", {
                            required: "Required",
                        })}
                        type="password"
                        placeholder="Confirm Password"
                        _placeholder={{ color: "gray" }}
                        border={
                            passwordMismatch
                                ? "1px solid red"
                                : ".5px solid gray"
                        }
                    />
                    {passwordMismatch && (
                        <Text color="red.500">Passwords do not match</Text>
                    )}
                </div>

                {/* phone number input */}
                <div className="mb-2">
                    <div className="font-light mb-1 text-[#AFAFAF]">Phone Number</div>
                    <PhoneInput
                        placeholder="Enter phone number"
                        value={phoneNumber}
                        onChange={(value) => {
                            if (value) {
                                setPhoneNumber(value);
                            }
                        }}
                        defaultCountry="SG"
                        className="text-black customPhoneInput"
                        international
                        withCountryCallingCode
                    />
                </div>

                <Button
                    onClick={handleSubmit(onSubmit)}
                    colorScheme="blue"
                    w="full"
                >
                    Create Account
                </Button>
            </div>
        </Box>
    );
};

export default SignUpStep;
