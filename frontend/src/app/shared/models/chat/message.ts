export interface Message {
  userId: string;
  userTName: string;
  userName: string;
  questionId: number;
  requestTime: string;
  text: string;
  answer?: string;
}
