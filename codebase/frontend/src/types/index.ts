export interface User {
  id: string;
  email: string;
  role: 'admin' | 'psicologo';
  name: string;
  crp?: string;
  specialty?: string;
  phone?: string;
  base_fee: number;
  package_sessions?: number;
  package_fee?: number;
  created_at: string;
  updated_at: string;
}

export interface Patient {
  id: string;
  psychologist_id: string;
  name: string;
  phone: string;
  birthdate?: string;
  age?: number;
  profession?: string;
  company?: string;
  city?: string;
  state?: string;
  marital_status?: string;
  consultation_fee: number;
  active: boolean;
  created_at: string;
  updated_at: string;
}

export interface FirstAnalysis {
  patient_id: string;
  main_complaint: string;
  symptom_diagnosis: string;
  developmental_influence: string;
  situational_issues: string;
  biological_factors: string;
  strengths_resources: string;
  addictions: string;
  stimuli: string;
  thoughts: string;
  behaviors: string;
  affects: string;
  physiological: string;
  treatment_goals: string;
  treatment_plan: string;
  updated_at: string;
}

export interface Session {
  id: string;
  patient_id: string;
  session_date: string;
  notes: string;
  status: 'pago' | 'pendente';
  meet_link?: string;
  outlook_event_id?: string;
  created_at: string;
  updated_at: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}
