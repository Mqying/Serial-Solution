import buildConfig from '../config/global';

export default function access(initialState) {
  const { currentUser, config } = initialState ?? {};
  const included = 1;

  if (buildConfig.locale == 'zh-CN') {
    return {
      canAdmin: currentUser && currentUser.access === 'admin',
      dielectric: (config.dielectron == included),
      water: (config.water == included),
      acid: (config.acid == included),
      flash: (config.flash == included),
      introduction: (config.logo == included),
    };
  }

  if (buildConfig.locale == 'en-US') {
    return {
      canAdmin: currentUser && currentUser.access === 'admin',
      dielectric: (config.dielectron == included),
      water: (config.water == included),
      acid: (config.acid == included),
      flash: (config.flash == included),
      introduction: false,
    };
  }
}
