import OSS from 'ali-oss'
import {} from 'lodash';

const ossclient = new OSS({
  region: 'oss-cn-hangzhou',
  accessKeyId: '',
  accessKeySecret: '',
  stsToken: '',
  bucket: '',
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
