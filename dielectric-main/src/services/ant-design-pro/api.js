// @ts-ignore

/* eslint-disable */
import { request } from 'umi'
/** get current user GET /api/currentUser */

export async function currentUser(options)
{
  return request('/api/currentUser', {
    method: 'GET',
    ...(options || {}),
    skipErrorHandler: true,
  })
}