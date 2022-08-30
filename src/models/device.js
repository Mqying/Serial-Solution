import { frontPage, nextPage, previousPage } from '../services/device/index.js';
export default {
  namespace: 'device',
  state: {
    detection_time: 'N/A',
    average: 0.0,
    items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    number: 0,
    quantity: 0,
    ratio1: 0,
    ratio2: 0,
    sampleNumber: 0,
    preTemperature: 0,
    pressure: 0,
    pointTemperature: 0,
    status: 204,
    errorCode: 13,
  },
  effects: {
    *frontPage({ payload }, { call, put }) {
      const response = yield call(frontPage, payload);
      yield put({
        type: 'saveDetectionData',
        payload: response,
      });
    },

    *nextPage({ payload }, { call, put }) {
      const response = yield call(nextPage, payload);
      yield put({
        type: 'saveDetectionData',
        payload: response,
      });
    },

    *previousPage({ payload }, { call, put }) {
      const response = yield call(previousPage, payload);
      yield put({
        type: 'saveDetectionData',
        payload: response,
      });
    },
  },

  reducers: {
    saveDetectionData(state, { payload }) {
      return {
        ...state,
        errorCode: payload.errorCode,
        status: payload.status,
        detectionTime: (payload.record === undefined ? 'N/A' : payload.record.detection_time),
        number: (payload.record === undefined ? 0 : payload.record.index),
        average: (payload.record === undefined ? 0 : payload.record.average),
        items: (payload.record === undefined ? [0, 0, 0, 0, 0, 0, 0, 0, 0, 0] : payload.record.items),
        quantity: (payload.record === undefined ? 0 : payload.record.quantity),
        ratio1: (payload.record === undefined ? 0 : payload.record.ratio1),
        ratio2: (payload.record === undefined ? 0 : payload.record.ratio2),
        sampleNumber: (payload.record === undefined ? 0 : payload.record.sample_number),
        preTemperature: (payload.record === undefined ? 0 : payload.record.pre_temperature),
        pointTemperature: (payload.record === undefined ? 0 : payload.record.point_temperature),
        pressure: (payload.record === undefined ? 0 : payload.record.pressure),
      };
    },
  },
};
