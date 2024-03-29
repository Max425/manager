import {Question} from "./question";

export interface Theme {
  id: number;
  label: string;
  categoryId?: number;
  children?: Question[];
}
