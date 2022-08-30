import { TableType } from '@/tableType.js';
import moment from 'moment';
import { getAllRecords } from '../services/history/index.js';
export default {
  namespace: 'history',
  state: {
    records: [{
      detection_time: 0,
      average: 0,
      index: 0,
      quantity: 0,
      ratio2: 0,
      ratio1: 0,
      sample_number: 0,
      pressure: 0,
      pre_temperature: 0,
      point_temperature: 0,
      items: [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
    }],
    status: 204,
    errorCode: 13,
  },
  effects: {
    *getAllRecords({ payload }, { call, put }) {
      const response = yield call(getAllRecords, payload);
      const dateFormat = 'YYYY-MM-DD HH:mm'

      switch (payload.type) {
        case TableType.dielectric:
          for (let id in response.records) {
            let row = response.records[id]

            row.average += ' kV';
            row.detectionTime = (row.detection_time == 0 ? 'N/A' : moment(row.detection_time).format(
              dateFormat
            ));

            for (let i in row.items) {
              let order = parseInt(i) + 1;
              row['item ' + order] = row.items[i] + ' kV';
            }

            row.id = id;
          }

        case TableType.water:
          for (let id in response.records) {
            let row = response.records[id]

            row.detectionTime = (row.detection_time == 0 ? 'N/A' : moment(row.detection_time).format(
              dateFormat
            ));
            row.quantity = row.quantity + " ug";
            row.ratio1 = (row.ratio1 === undefined) ? 'N/A' : row.ratio1 + ' PPM';
            row.ratio2 = (row.ratio2 === undefined) ? 'N/A' : row.ratio2 + ' %';
            row.id = id;
          }

        case TableType.acid:
          for (let id in response.records) {
            let row = response.records[id]

            row.detectionTime = (row.detection_time == 0 ? 'N/A' : moment(row.detection_time).format(
              dateFormat
            ));
            for (let i in row.items) {
              let order = parseInt(i) + 1;
              row['item' + order] = row.items[i];
            }
            row.id = id;
          }

        case TableType.flash:
          for (let id in response.records) {
            let row = response.records[id]

            row.detectionTime = (row.detection_time == 0 ? 'N/A' : moment(row.detection_time).format(
              dateFormat
            ));
            row.sampleNumber = row.sample_number;
            row.preTemperature = row.pre_temperature;
            row.pointTemperature = row.point_temperature;
            row.pressure = row.pressure;
            row.id = id;
          }
      }

      yield put({
        type: 'saveRecords',
        payload: response,
      });
    },
  },
  reducers: {
    saveRecords(state, { payload }) {
      return {
        ...state,
        status: payload.status,
        errorCode: payload.errorCode,
        records: payload.records,
      };
    },
  },
};
