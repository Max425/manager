import {Theme} from "./theme";

export interface Category {
  id: number;
  label: string;
  children?: Theme[];
}
