export interface Chat {
  userId: string;
  userTName: string;
  userName: string;
  countQuestion: number;
  isHidden: boolean;
  questions?: ChatQuestion[];
}

export  interface ChatQuestion {
  questionId: number;
  requestTime: string;
  text: string;
}
