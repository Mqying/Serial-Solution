/* eslint-disable react/jsx-key */
/* eslint-disable react-hooks/exhaustive-deps */
import { Descriptions, Spin } from 'antd';
import { connect } from 'dva';
import moment from 'moment';
import { useIntl } from 'umi';
import { useEffect, useState } from 'react';
import { QueryPageFramework } from '../../components/QueryPageFramework/index.jsx';
import { TableType } from '@/tableType.js';
import styles from '../../global.less';
import { global } from '../../global';

const DielectricQuery = ({ dispatch, data, loading }) => {
  const intl = useIntl()
  const [pageSize, setPageSize] = useState(window.innerHeight);
  useEffect(() => {
    if (!global.isFrozen() && global.isFirst()) {
      dispatch({
        type: 'device/frontPage',
        payload: {
          type: TableType.dielectric,
        }
      });
    }
    window.addEventListener('resize', resizeUpdate);

    return () => {
      window.removeEventListener('resize', resizeUpdate);
    }
  }, [1]);

  const resizeUpdate = (e) => {
    setPageSize(e.target.innerHeight);
  }

  const itemCount = 7.5;
  const subtractionHeight = 204;
  const descriptionItemHeight = (pageSize - subtractionHeight) / itemCount;
  const successCode = 200;

  const stateCardStyle = () => {
    if (data.status === successCode) {
      return styles.stateCodeCardSuccess
    } else {
      return styles.stateCodeCardFail
    }
  }

  return (
    (data.items ?
      <QueryPageFramework
        dispatch={dispatch}
        loading={loading}
        type={TableType.dielectric}
      >
        <div className={stateCardStyle()}>
          {
            data.status === successCode ?
              intl.formatMessage({ id: "pages.table.requestSuccess" }) :
              intl.formatMessage({ id: "pages.table.requestFailureCode" }) +
              data.errorCode
          }
        </div>
        <Descriptions bordered className={styles.table}>
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={intl.formatMessage({ id: "pages.table.detectionTime" })} span={4}>
            {data.detectionTime == 'N/A' ? data.detectionTime : moment(data.detectionTime).format('YYYY-MM-DD HH:mm')}
          </Descriptions.Item >
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={intl.formatMessage({ id: "pages.table.number" })} span={2}>
            {data.number}
          </Descriptions.Item >
          <Descriptions.Item style={{ height: descriptionItemHeight }}
            label={intl.formatMessage({ id: "pages.table.avg" })} span={4}>
            {data.average + ' kV'}
          </Descriptions.Item>

          {
            data.items.map((item, index) => {
              return (
                <Descriptions.Item key={index} style={{ height: descriptionItemHeight }}
                  label={intl.formatMessage({ id: "pages.table.frequency" }) + (index + 1)} span={2}>
                  {item + ' kV'}
                </Descriptions.Item>
              )
            })
          }
        </Descriptions>
      </QueryPageFramework > : <Spin size="large" />)

  );
};

export default connect(({ device, loading }) => ({
  data: device,
  loading: loading.models.device,
}))(DielectricQuery);