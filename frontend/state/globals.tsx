import { atom } from "jotai";

export const userAtom = atom<{ email?: string, userId?: string} | null>(null);
export const authModalAtom = atom<boolean>(false);