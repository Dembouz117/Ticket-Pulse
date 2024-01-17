//mock section data
import { sectionsSchema, sectionAlgo } from "@/types/index";

const sections: sectionAlgo[] = [
    { id: 1, section: "A", category:"CAT3", seats: 10, seatsAvailable: 10 },
    { id: 2, section: "B", category:"CAT3", seats: 15, seatsAvailable: 15 },
    { id: 3, section: "A", category:"CAT2", seats: 10, seatsAvailable: 10 },
    { id: 4, section: "B", category:"CAT2", seats: 15, seatsAvailable: 15 },
    { id: 5, section: "A", category:"CAT1", seats: 10, seatsAvailable: 10 },
    { id: 6, section: "B", category:"CAT1", seats: 15, seatsAvailable: 15 },
];

export { sections };