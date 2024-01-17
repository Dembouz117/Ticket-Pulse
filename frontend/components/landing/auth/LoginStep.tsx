import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Box, Button, Input, useToast, Text, VStack } from "@chakra-ui/react";
import "react-phone-number-input/style.css";
import PhoneInput, { isValidPhoneNumber } from "react-phone-number-input";
import axios from "axios";
import { getAuthApiUrl } from "@/utilities/common";
import { get } from "http";

interface LoginFormData {
    email: string;
    password: string;
}

interface LoginStepProps {
    setStep: (step: string) => void;
    setEmail: (email: string) => void;
    setPassword: (password: string) => void;
}

const LoginStep: React.FC<LoginStepProps> = ({
    setStep,
    setEmail,
    setPassword,
}) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<LoginFormData>();
    const toast = useToast();

    const [phoneNumber, setPhoneNumber] = useState<string>("");

    const onSubmit = async (data: LoginFormData) => {
        try {
            console.log("Before get auth api url");
            console.log(process.env.REACT_APP_AUTH_API_URL);
            console.log(getAuthApiUrl());
            const response = await axios.post(`${getAuthApiUrl()}/login`, {
                email: data.email,
                password: data.password,
            });

            console.log(data);

            if (response.status === 200) {
                setEmail(data.email);
                setPassword(data.password);
                setStep("twoFactor");

                //currently broken and just goes into catch block
            } else if (response.status === 401) {
                toast({
                    title: "Login failed",
                    description: "Invalid email or password.",
                    status: "error",
                    duration: 3000,
                    isClosable: true,
                });
            }
        } catch (error) {
            toast({
                title: "Error",
                description: "An error occurred while logging in.",
                status: "error",
                duration: 3000,
                isClosable: true,
            });
        }
    };

    return (
        <Box w="full" p={4}>
            <div className="flex flex-col align-middle justify-start gap-4 p-3">
                <div className="text-3xl mb-2">Welcome Back!</div>

                <div className="">
                    <div className="font-light mb-1 text-[#AFAFAF]">Email</div>
                    <Input
                        {...register("email", {
                            required: "Required",
                            pattern: {
                                value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
                                message: "Invalid email address",
                            },
                        })}
                        placeholder="johndoe@example.com"
                        _placeholder={{ color: "gray" }}
                        border={
                            errors.email ? "1px solid red" : ".5px solid gray"
                        }
                        data-cy="email"
                    />
                    {errors.email && (
                        <Text color="red.500">{errors.email.message}</Text>
                    )}
                </div>

                <div>
                    <div className="font-light mb-1 text-[#AFAFAF]">
                        Password
                    </div>
                    <Input
                        {...register("password", {
                            required: "Required",
                        })}
                        type="password"
                        autoComplete="off"
                        placeholder="Password"
                        _placeholder={{ color: "gray" }}
                        border={
                            errors.email ? "1px solid red" : ".5px solid gray"
                        }
                        data-cy="password"
                    />
                    {errors.password && (
                        <Text color="red.500">{errors.password.message}</Text>
                    )}
                </div>

                <button
                    onClick={handleSubmit(onSubmit)}
                    className="bg-button hover:bg-button-hover text-white rounded-lg w-full py-2"
                    data-cy="login"
                >
                    Log In
                </button>

                <div className="flex flex-row gap-2 mb-2">
                    <div className="text-neutral-500">
                        Don't have an account?{" "}
                    </div>
                    <Button
                        variant="link"
                        color="blue.500"
                        onClick={() => setStep("signUp")}
                    >
                        Sign Up
                    </Button>
                </div>
            </div>
        </Box>
    );
};

export default LoginStep;
