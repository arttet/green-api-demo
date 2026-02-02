import { dev } from '$app/environment';
import { env } from '$env/dynamic/public';

export interface AppConfig {
  readonly apiBaseUrl: string;
  readonly isProduction: boolean;
  readonly timeout: number;
}

const DEFAULT_API_URL = 'http://localhost:8080/v1/api/proxy' as const;

export const config: AppConfig = {
  apiBaseUrl: env.PUBLIC_API_BASE_URL?.trim() || DEFAULT_API_URL,
  isProduction: !dev,
  timeout: 5000,
} as const;
