import OSS from 'ali-oss'
import {} from 'lodash';

console.log(process.env)

const ossclient = new OSS({
  region: process.env.REACT_APP_OSS_REGION,
  accessKeyId: process.env.REACT_APP_OSS_ACCESSKEYID,
  accessKeySecret: process.env.REACT_APP_OSS_ACCESSKEYSECRET,
  bucket: process.env.REACT_APP_OSS_BUCKET,
});

function buildFileName(path, postfix) {
  return path + new Date().getTime() + postfix;
}

function OSSUploaderFile(file, path, onSucess, onError) {
  console.error(file);
  const index = file.name.lastIndexOf('.');

  if (index === -1) {
    return onError('file error.');
  }
  const postfix = file.name.substr(index);

  ossclient.put(buildFileName(path, postfix), file)
    .then((ret) => {
      onSucess(ret, file);
    })
    .catch(onError);
}

export default OSSUploaderFile;
