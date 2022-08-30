/* eslint-disable react-hooks/exhaustive-deps */
import { Descriptions } from 'antd';
import { connect } from 'dva';
import moment from 'moment';
import { useIntl } from 'umi';
import { useEffect, useState } from 'react';
import { QueryPageFramework } from '../../components/QueryPageFramework/index.jsx';
import { TableType } from '@/tableType.js';
import styles from '../../global.less';
import { global } from '../../global';

const WaterQuery = ({ dispatch, data, loading }) => {
  const intl = useIntl()
  const [pageSize, setPageSize] = useState(window.innerHeight);
  useEffect(() => {
    if (!global.isFrozen() && global.isFirst()) {
      dispatch({
        type: 'device/frontPage',
        payload: {
          type: TableType.water,
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

  const itemCount = 5.6;
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
    <QueryPageFramework
      dispatch={dispatch}
      loading={loading}
      type={TableType.water} description={descriptionItemHeight}
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
          {data.detectionTime == 'N/A' ? data.detectionTime
            : moment(data.detectionTime).format('YYYY-MM-DD hh:mm')}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.number" })} span={4}>
          {data.number}
        </Descriptions.Item >
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.quantity" })} span={4}>
          {data.quantity + ' ug'}
        </Descriptions.Item>
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.percentage" })} span={4}>
          {data.ratio1 === undefined ? "N/A" : data.ratio1 + ' PPM'}
        </Descriptions.Item>
        <Descriptions.Item style={{ height: descriptionItemHeight }}
          label={intl.formatMessage({ id: "pages.table.percentage" })} span={4}>
          {data.ratio2 === undefined ? "N/A" : data.ratio2 + ' %'}
        </Descriptions.Item>
      </Descriptions>
    </QueryPageFramework >

  );
};

export default connect(({ device, loading }) => ({
  data: device,
  loading: loading.models.device,
}))(WaterQuery);