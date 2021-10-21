const { spawn } = require('child_process');

const vlcPath = '/Applications/VLC.app/Contents/MacOS/VLC';

const play = (input, args) => {
  const vlc = spawn(vlcPath, [input].concat(args));
  vlc.stdout.on('data', d => console.log(d.toString().trim()));
  vlc.stderr.on('data', e => console.error(e.toString().trim()));
  return vlc;
};

module.exports = { play };
