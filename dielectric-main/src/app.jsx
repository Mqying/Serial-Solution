// import { Footer } from '@/components/Footer';
import RightContent from '@/components/RightContent';
import { BookOutlined, LinkOutlined } from '@ant-design/icons';
import { PageLoading } from '@ant-design/pro-layout';
import { history, Link } from 'umi';
import defaultSettings from '../config/defaultSettings';
import { currentUser as queryCurrentUser } from './services/ant-design-pro/api';
import { getConfig } from './services/config'
import buildConfig from '../config/global';
import zhCN from '../src/locales/zh-CN/menu'
import enUS from '../src/locales/en-US/menu'

const local = buildConfig.locale == 'zh-CN' ? zhCN : enUS
const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';

export const initialStateConfig = {
  loading: <PageLoading />,
};

export async function getInitialState() {
  const fetchUserInfo = async () => {
    try {
      const msg = await queryCurrentUser();
      return msg.data;
    } catch (error) {
      history.push(loginPath);
    }

    return undefined;
  };

  const response = await getConfig();

  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
      settings: defaultSettings,
      config: response.config,
    };
  }

  return {
    fetchUserInfo,
    settings: defaultSettings,
    config: response.config,
  };
}

export const layout = ({ initialState, setInitialState }) => {
  const included = 1;
  let title;

  if (initialState.config.dielectron == included) {
    title = local["menu.title.dielectric"];
  } else if (initialState.config.water == included) {
    title = local["menu.title.water"];
  } else if (initialState.config.acid == included) {
    title = local["menu.title.acid"];
  } else if (initialState.config.flash == included) {
    title = local["menu.title.flash"];
  } else {
    title = local["menu.title.error"];
  }

  return {
    title: title,
    logo: initialState.config.logo == included ? './logo(1).svg' : false,
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    waterMarkProps: {
      //content: initialState?.currentUser?.name,
    },
    // footerRender: () => <Footer />,
    // onPageChange: () => {
    //   const { location } = history; // 如果没有登录，重定向到 login

    //   if (!initialState?.currentUser && location.pathname !== loginPath) {
    //     history.push(loginPath);
    //   }
    // },
    links: isDev
      ? [
        <Link key="openapi" to="/umi/plugin/openapi" target="_blank">
          <LinkOutlined />
          <span>OpenAPI 文档</span>
        </Link>,
        <Link to="/~docs" key="docs">
          <BookOutlined />
          <span>业务组件文档</span>
        </Link>,
      ]
      : [],
    menuHeaderRender: undefined,
    fixedHeader: true,
    menuHeaderRender: (logo, title) => {
      if (logo == false) {
        return (
          <>
            <a style={{ minHeight: '100%', margin: 0, padding: 0 }}>
              <h1>{title}</h1>
            </a>
          </>
        )
      } else {
        return (
          <>
            <a style={{ minHeight: '100%', margin: 0, padding: 0 }}>
              <img src="./logo(1).svg" style={{ height: '44px' }} alt="logo" /><h1>{title}</h1>
            </a>
          </>
        )
      }
    },
    rightContentRender: () => {
      return (
        <></>
      )
    },
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children, props) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          {!props.location?.pathname?.includes('/login')}
        </>
      );
    },
    ...initialState?.settings,
  };
};