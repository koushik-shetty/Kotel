"use strict";

var gulp = require("gulp"),
	chalk = require("chalk"),
	childProcess = require("child_process"),
	exec = childProcess.exec,
	gulp = require('gulp'),
	concat = require('gulp-concat'),
	uglify = require('gulp-uglify'),
	react = require('gulp-react'),
	htmlreplace = require('gulp-html-replace'),
	outDir = "out",
	exe = "TWLibrary",

	paths = {
		out: "out",
		jsOut:"out/public/javascript",
		cssOut: ""
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

gulp.task("compile-go",function(callback){
	execute("go build -o " + outDir +"/"+exe,callback);
});


gulp.task("build",["compile-go"]);

gulp.task("default",["format","build"]);

gulp.task("run",function(callback){
	execute("./"+outDir+"/"+exe + "&",callback);
});