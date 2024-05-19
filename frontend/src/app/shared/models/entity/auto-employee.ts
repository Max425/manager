import {Employee} from "./employee";

export interface AutoEmployee {
  position: string;
  employee: Employee | undefined;
  pin: boolean;
}
