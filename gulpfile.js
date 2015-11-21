'use strict';

var gulp = require('gulp'),
	chalk = require('chalk'),
	childProcess = require('child_process'),
	exec = childProcess.exec,
	gulp = require('gulp'),
	concat = require('gulp-concat'),
	uglify = require('gulp-uglify'),
	babelify = require('babelify'),
	babel = require('gulp-babel'),
	browserify = require('browserify'),
	source = require('vinyl-source-stream'),
	react = require('gulp-react'),
	jshint = require('gulp-jshint'),
	exe = 'TWLibrary.exe',
	outDir = 'out',
	jsMainfile = 'main.js',
	paths = {
		Target: outDir + '/' + exe,
		jsSrc: 'javascript/src',
		jsOut: outDir + '/public/javascript',
		cssSrc: ['views/assets/images/*.*','views/assets/*.css','views/assets/**/*.css'],
		cssOut: outDir + '/public/styles',
	},
	jsEntry= paths.jsSrc + '/' + jsMainfile;


var execute = function(command, callback) {
    if (!callback) {
      throw 'ArgumentMissing when running ' + command + ' (callback: ' + callback + ')';
    }
    console.log(' %s: %s', chalk.grey('>> Running'), chalk.cyan(command));
    exec(command, function(err, stdout, stderr) {
      if (stdout !== '') {
        console.log(' >> %s:\n%s', chalk.cyan(command), chalk.grey(stdout));
      }
      if (stderr !== '') {
        console.error(' >> %s:\n%s', chalk.cyan(command), chalk.red(stderr));
      }
      callback(err);
      return stdout;
    });
};


gulp.task('format',function(callback){
	execute('go fmt ./...',callback);	
});

gulp.task('jshint',function(callback){
	execute('jshint --verbose ' + paths.jsSrc,callback);
});

/**********************************Build tasks*************************************/

gulp.task('build-js',['build-assets'],function(callback){
	return browserify({
		entries: jsEntry,
		debug: true
	})
	.transform(babelify,{presets : ['es2015','react']})
	.bundle()
	.pipe(source(jsMainfile))
	.pipe(gulp.dest(paths.jsOut));
});

gulp.task('build-assets',function(callback){
	gulp.src(paths.cssSrc)
	.pipe(gulp.dest(paths.cssOut))
});


gulp.task('compile-go',function(callback){
	execute('go build -o ' + paths.Target, callback);
});


gulp.task('build',['compile-go','build-js']);

gulp.task('clean',function(callback){
	execute('rm -rf out',callback);	
});

/**********************************************************************************/

gulp.task('default',['format','build']);

gulp.task('run',function(callback){
	execute('./'+ paths.Target + ' &', callback);
});