import {
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalOverlay,
} from "@chakra-ui/react";
import { useState } from "react";
import SignUpStep from "./SignUpStep";
import LoginStep from "./LoginStep";
import TwoFactorStep from "./TwoFactorStep";

interface AuthModalProps {
  open: boolean;
  onClose: () => void;
}

const AuthModal = ({ open, onClose }: AuthModalProps) => {
  const [step, setStep] = useState<string>("login");
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleClose = () => {
    // reset state to login
    setStep("login");
    onClose();
  };

  return (
    // <div data-cy="auth-modal">
      <Modal
        isOpen={open}
        onClose={handleClose}
        closeOnOverlayClick={false}
        size={"lg"}
        data-cy="auth-modal"
      >
        <ModalOverlay />
        <ModalContent bg={"#1F1F1F"} textColor={"white"}>
          <ModalCloseButton data-cy="auth-close"/>
          <ModalBody>
            {step === "signUp" && (
              <SignUpStep
                onClose={handleClose}
                setStep={setStep}
                setEmail={setEmail}
                setPassword={setPassword}
              />
            )}
            {step === "login" && (
              <LoginStep
                setStep={setStep}
                setEmail={setEmail}
                setPassword={setPassword}
              />
            )}
            {step === "twoFactor" && (
              <TwoFactorStep
                onClose={handleClose}
                setStep={setStep}
                email={email}
                password={password}
              />
            )}
          </ModalBody>
        </ModalContent>
      </Modal>
    // </div>
  );
};

export default AuthModal;
