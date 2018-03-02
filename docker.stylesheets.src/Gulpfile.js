'use strict';

const gulp = require('gulp');
const sass = require('gulp-sass');

gulp.task('default', () => {
  return gulp.src('src/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest('dist/'));
});

gulp.task('sass:watch', () => {
  gulp.watch('src/*.scss', ['default']);
});
