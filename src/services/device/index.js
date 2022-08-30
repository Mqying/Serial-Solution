import { request } from 'umi'
import { global } from '../../global'

export async function frontPage(params) {
  const successCode = 200;
  global.setFrozen();

  let response = {
    status: 500,
    errorCode: 12,
    record: {
      detection_time: 'N/A',
      average: 0.0,
      items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
      index: 0,
      quantity: 0,
      ratio1: 0,
      ratio2: 0,
      sample_number: 0,
      pre_temperature: 0,
      point_temperature: 0,
      pressure: 0
    }
  }

  try {
    response = await request('/api/v1/device/frontPage', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
      params: params,
    })
  } catch (err) {
  }

  if (response.status === successCode) {
    global.setFirst();
  }

  global.clearFrozen();

  return response
}

export async function previousPage(params) {
  global.setFrozen();

  let response = {
    status: 500,
    errorCode: 12,
    record: {
      detection_time: 'N/A',
      average: 0.0,
      items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
      index: 0,
      quantity: 0,
      ratio1: 0,
      ratio2: 0,
      sample_number: 0,
      pre_temperature: 0,
      point_temperature: 0,
      pressure: 0
    }
  }

  try {
    response = await request('/api/v1/device/previousPage', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
      params: params,
    })
  } catch (err) {
  }

  global.clearFrozen();

  return response
}

export async function nextPage(params) {
  global.setFrozen();

  let response = {
    status: 500,
    errorCode: 12,
    record: {
      detection_time: 'N/A',
      average: 0.0,
      items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
      index: 0,
      quantity: 0,
      ratio1: 0,
      ratio2: 0,
      sample_number: 0,
      pre_temperature: 0,
      point_temperature: 0,
      pressure: 0
    }
  }

  try {
    response = await request('/api/v1/device/nextPage', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
      params: params,
    })
  } catch (err) {
  }

  global.clearFrozen();

  return response
}

export async function print() {
  global.setFrozen();

  let response = {
    status: 500,
    errorCode: 12,
  }

  try {
    response = await request('/api/v1/device/print', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
    })
  } catch (err) {
  }

  global.clearFrozen();

  return response
}
