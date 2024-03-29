export interface Question {
  id: number;
  label: string;
  answer: string;
  themeId?: number;
  isFrequentQuestion?: boolean;
  showQuestion?: boolean;
  dateExpire?: Date;
}
