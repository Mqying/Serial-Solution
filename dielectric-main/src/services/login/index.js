import { request } from 'umi';
export async function login() {
  let response = {
    status: 500
  }

  try {
    response = await request('/api/v1/admin/login', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
      skipErrorHandler: true,
    });
  } catch (error) {

  }

  return response;
}
