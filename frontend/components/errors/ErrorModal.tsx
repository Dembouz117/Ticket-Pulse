import { Modal, ModalCloseButton, ModalContent, ModalOverlay } from "@chakra-ui/react";



interface ErrorModalProps{
    settings: {
        visible: boolean,
        message: string
    },
    modalHandler: any
}
export const ErrorModal = ({settings, modalHandler}: ErrorModalProps) => {
  
    return (
        <>
        {settings.visible && 
        <Modal
        isOpen={true}
        onClose={modalHandler}
        closeOnOverlayClick={true}
        size={"lg"}
    >
        <ModalOverlay />
        <ModalContent bg={"#1F1F1F"} textColor={"white"} className = "py-12 px-4 h-auto">
            <ModalCloseButton />
            <div className = "flex-col space-y-6">
                <h2 className = "text-3xl text-red-600">Unfortunately, something went wrong.</h2>
                <h4>{settings.message}</h4>
            </div>

        </ModalContent>
    </Modal>
        }
        {!settings.visible && 
                <Modal
                isOpen={true}
                onClose={modalHandler}
                closeOnOverlayClick={true}
                size={"lg"}
            >
                <ModalOverlay />
                <ModalContent bg={"#1F1F1F"} textColor={"white"} className = "py-12 px-4 h-auto">
                    <ModalCloseButton />
                    <div className = "flex-col space-y-6">
                        <h2 className = "text-3xl">There is no error.</h2>
                    </div>
        
                </ModalContent>
            </Modal>
        }
        </>
        
    )
}

export default ErrorModal;