/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: apisupport@spyderbat.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.SpyderbatApi);
  }
}(this, function(expect, SpyderbatApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new SpyderbatApi.SourceApi();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('SourceApi', function() {
    describe('integrationSoarSrcList', function() {
      it('should call integrationSoarSrcList successfully', function(done) {
        //uncomment below and update the code to test integrationSoarSrcList
        //instance.integrationSoarSrcList(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('srcCreate', function() {
      it('should call srcCreate successfully', function(done) {
        //uncomment below and update the code to test srcCreate
        //instance.srcCreate(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('srcDelete', function() {
      it('should call srcDelete successfully', function(done) {
        //uncomment below and update the code to test srcDelete
        //instance.srcDelete(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('srcList', function() {
      it('should call srcList successfully', function(done) {
        //uncomment below and update the code to test srcList
        //instance.srcList(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('srcLoad', function() {
      it('should call srcLoad successfully', function(done) {
        //uncomment below and update the code to test srcLoad
        //instance.srcLoad(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('srcUpdate', function() {
      it('should call srcUpdate successfully', function(done) {
        //uncomment below and update the code to test srcUpdate
        //instance.srcUpdate(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
