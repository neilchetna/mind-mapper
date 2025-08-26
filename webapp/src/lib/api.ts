import { PUBLIC_BACKEND_URL } from '$env/static/public';

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';

interface RequestConfig extends Omit<RequestInit, 'body'> {
	params?: Record<string, string | number | boolean>;
	body?: unknown;
	headers?: Record<string, string>;
}

interface ApiResponse<T> {
	data: T;
	status: number;
	ok: boolean;
}

export class ApiClient {
	private baseURL: string;
	private authToken?: string;

	constructor() {
		this.baseURL = PUBLIC_BACKEND_URL;
	}

	public newAuthToken(token: string) {
		this.authToken = token;
	}

	private buildUrl(path: string, params?: RequestConfig['params']): string {
		const url = new URL(`${this.baseURL}${path}`);
		if (params) {
			Object.entries(params).forEach(([key, value]) => url.searchParams.append(key, String(value)));
		}
		return url.toString();
	}

	private async request<T>(
		method: HttpMethod,
		path: string,
		config: RequestConfig = {}
	): Promise<ApiResponse<T>> {
		const url = this.buildUrl(path, config.params);

		const headers: HeadersInit = {
			Authorization: `Bearer ${this.authToken}`,
			'Content-Type': 'application/json',
			...(config.headers || {})
		};

		const res = await fetch(url, {
			...config,
			method,
			headers,
			body: config.body ? JSON.stringify(config.body) : undefined
		});

		const contentType = res.headers.get('content-type');
		const data = contentType?.includes('application/json') ? await res.json() : await res.text();

		if (!res.ok) {
			throw new Error(`API Error: ${res.status} ${JSON.stringify(data)}`);
		}

		return {
			data: data as T,
			status: res.status,
			ok: res.ok
		};
	}

	get<T>(path: string, config?: RequestConfig) {
		return this.request<T>('GET', path, config);
	}

	post<B, R>(path: string, body?: B, config?: RequestConfig) {
		return this.request<R>('POST', path, { ...config, body });
	}

	put<B, R>(path: string, body?: B, config?: RequestConfig) {
		return this.request<R>('PUT', path, { ...config, body });
	}

	patch<B, R>(path: string, body?: B, config?: RequestConfig) {
		return this.request<R>('PATCH', path, { ...config, body });
	}

	delete<T>(path: string, config?: RequestConfig) {
		return this.request<T>('DELETE', path, config);
	}
}
