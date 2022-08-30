/* eslint-disable react/jsx-key */
/* eslint-disable react/jsx-no-undef */
import { PageLoading } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { Button, Card, message } from 'antd';
import { connect } from 'dva';
import { useEffect, useRef, useState } from 'react';
import { useIntl, useModel } from 'umi';
import { exportRecords } from "../../services/exportRecords/index"
import { TableType } from '@/tableType.js';
import styles from '../../global.less';
import { global } from '../../global';

const Record = (props) => {
  const intl = useIntl();
  const { initialState } = useModel('@@initialState');
  const { dispatch, data, loading } = props;
  const actionRef = useRef();
  const [selectedRowsState, setSelectedRows] = useState([]);
  const device = Object.keys(initialState.config)
  let type = '0';
  const successCode = 200;

  for (let i = 0; i < Object.keys(TableType).length; i++) {
    if (initialState.config[device[i]] == 1) {
      type = (i + '')
      break
    }
  }

  useEffect(() => {
    if (!global.isFrozen()) {
      dispatch({
        type: 'history/getAllRecords',
        payload: {
          type: type,
        }
      });
    }
  }, [type]);

  const getTypeColumns = (type) => {
    switch (type) {
      case TableType.dielectric:
        let dielectricColumns = [
          {
            title: intl.formatMessage({ id: "pages.table.detectionTime" }),
            dataIndex: 'detectionTime',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.export.head.index" }),
            dataIndex: 'index',
            width: "7.5%",
          },
          {
            title: intl.formatMessage({ id: "pages.table.avg" }),
            dataIndex: 'average',
            width: "7.5%",
          },
        ];

        for (let i = 1; i <= 10; i++) {
          dielectricColumns.push({
            title: intl.formatMessage({ id: "pages.table.frequency" }) + i,
            dataIndex: 'item ' + i,
            width: "7.5%",
          });
        }

        return dielectricColumns;

      case TableType.water:
        let waterColumns = [
          {
            title: intl.formatMessage({ id: "pages.table.detectionTime" }),
            dataIndex: 'detectionTime',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.quantity" }),
            dataIndex: 'quantity',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.percentage" }),
            dataIndex: 'ratio1',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.percentage" }),
            dataIndex: 'ratio2',
            fixed: "left",
          },
        ];

        return waterColumns;

      case TableType.acid:
        let acidColumns = [
          {
            title: intl.formatMessage({ id: "pages.table.detectionTime" }),
            dataIndex: 'detectionTime',
            fixed: "left",
          },
        ];

        for (let i = 1; i <= 6; i++) {
          acidColumns.push({
            title: "PH" + i,
            dataIndex: 'item' + i,
            width: "13%",
          });
        }

        return acidColumns;

      case TableType.flash:
        let flashColumns = [
          {
            title: intl.formatMessage({ id: "pages.table.detectionTime" }),
            dataIndex: 'detectionTime',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.atmosphericPressure" }),
            dataIndex: 'pressure',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.preFlashTemperature" }),
            dataIndex: 'preTemperature',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.flashPointTemperature" }),
            dataIndex: 'pointTemperature',
            fixed: "left",
          },
          {
            title: intl.formatMessage({ id: "pages.table.sampleSerialNumber" }),
            dataIndex: 'sampleNumber',
            fixed: "left",
          }
        ];

        return flashColumns;
    }
  }

  if (loading) {
    return <PageLoading />;
  }

  const stateCardStyle = () => {
    if (data.status === successCode) {
      return styles.stateCodeCardSuccess
    } else {
      return styles.stateCodeCardFail
    }
  }

  return (
    <Card className={styles.card} >
      <div className={stateCardStyle()}>
        {
          data.status === successCode ?
            intl.formatMessage({ id: "pages.table.requestSuccess" }) :
            intl.formatMessage({ id: "pages.table.requestFailureCode" }) +
            data.errorCode
        }
      </div>
      <ProTable
        actionRef={actionRef}
        rowKey="id"
        search={false}
        pagination={{
          showQuickJumper: true,
          size: 'middle',
          pageSize: 15,
        }}
        request={() => {
          return Promise.resolve({
            data: data.records,
            success: true,
          });
        }}
        columns={getTypeColumns(type)}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
        toolBarRender={() => [
          <Button
            key="button"
            type="primary"
            onClick={() => {
              if (selectedRowsState.length == 0) {
                message.error(intl.formatMessage({ id: "pages.table.noDataWasSelected" }));

                return;
              }

              exportRecords(selectedRowsState, type, intl.locale);
            }}
          >
            {intl.formatMessage({ id: "pages.table.exportAsExcel" })}
          </Button>,
        ]}
      />
    </Card >
  );
};

export default connect(({ history, loading }) => ({
  data: history,
  loading: loading.models.history,
}))(Record);