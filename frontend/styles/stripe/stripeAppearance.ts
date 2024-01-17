import { Appearance } from "@stripe/stripe-js";

//controls for all payment elements
export const appearance: Appearance = {
    theme: "night",
    labels: "floating",
    rules: {
        ".Input:hover:" : {
            border : "2px solid #EC4899",
            borderColor : "#EC4899"
        },
        ".Input:focus" : {
            border : "2px solid #EC4899",
            borderColor : "#EC4899" 
        }
    }
};