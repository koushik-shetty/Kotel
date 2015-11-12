"use strict";

var gulp = require("gulp"),
	chalk = require("chalk"),
	childProcess = require("child_process"),
	exec = childProcess.exec,
	gulp = require('gulp'),
	concat = require('gulp-concat'),
	uglify = require('gulp-uglify'),
	babel = require('babelify'),
	browserify = require('browserify'),
	react = require('gulp-react'),
	jslint = require('jslint'),
	exe = "TWLibrary",
	outDir = "out",
	paths = {
		Target: outDir + "/" + exe,
		jsSrc: ["javascript/*.js","javascript/**/*.js","javascript/**/**/*.js"],
		jsOut: outDir + "/public/javascript",
		cssSrc: ["views/styles/*.css","views/styles/**/*.css"],
		cssOut: outDir + "/public/styles",
	};

var execute = function(command, callback) {
    if (!callback) {
      throw "ArgumentMissing when running '" + command + "' (callback: " + callback + ")";
    }
    console.log(" %s: %s", chalk.grey(">> Running"), chalk.cyan(command));
    exec(command, function(err, stdout, stderr) {
      if (stdout !== "") {
        console.log(" >> %s:\n%s", chalk.cyan(command), chalk.grey(stdout));
      }
      if (stderr !== "") {
        console.error(" >> %s:\n%s", chalk.cyan(command), chalk.red(stderr));
      }
      callback(err);
      return stdout;
    });
};


gulp.task("format",function(callback){
	execute("go fmt ./...",callback);	
});

/**********************************Build tasks*************************************/

gulp.task("build-js",function(callback){
	gulp.src(paths.jsSrc)
		.pipe(babel())
		.pipe(browserify())
		.pipe(jslint())
		.pipe(gulp.dest(paths.jsOut))
})

gulp.task("compile-go",function(callback){
	execute("go build -o " + paths.Target, callback);
});


gulp.task("build",["compile-go","build-js"]);

/**********************************************************************************/

gulp.task("default",["format","build"]);

gulp.task("run",function(callback){
	execute("./"+ paths.Target + " &", callback);
});