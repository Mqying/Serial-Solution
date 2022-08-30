import { request } from 'umi';
import { global } from '../../global';

export async function getAllRecords(payload) {
  global.setFrozen();

  let response = {
    status: 500,
    errorCode: 12,
    records: [{
      detection_time: 0,
      average: 0,
      quantity: 0,
      ratio2: 0,
      ratio1: 0,
      sample_number: 0,
      pressure: 0,
      pre_temperature: 0,
      point_temperature: 0,
      items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    }]
  };

  try {
    response = await request('/api/v1/record/get', {
      method: 'GET',
      prefix: 'http://127.0.0.1:9090',
      params: payload,
    });
  } catch (err) {
  }

  global.clearFrozen();

  return response;
}
