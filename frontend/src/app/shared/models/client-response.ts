export interface ClientResponse<T> {
  status: number;
  message: string;
  payload: T | undefined;
}
