/**
 * Spyderbat API UI & Public APIs
 * Restful APIs for use by UI & customers.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
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
    factory(root.expect, root.Sbapi);
  }
}(this, function(expect, Sbapi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Sbapi.InvestigationApi();
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

  describe('InvestigationApi', function() {
    describe('investigationCreate', function() {
      it('should call investigationCreate successfully', function(done) {
        //uncomment below and update the code to test investigationCreate
        //instance.investigationCreate(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationDelete', function() {
      it('should call investigationDelete successfully', function(done) {
        //uncomment below and update the code to test investigationDelete
        //instance.investigationDelete(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationList', function() {
      it('should call investigationList successfully', function(done) {
        //uncomment below and update the code to test investigationList
        //instance.investigationList(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationListVersions', function() {
      it('should call investigationListVersions successfully', function(done) {
        //uncomment below and update the code to test investigationListVersions
        //instance.investigationListVersions(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationLoad', function() {
      it('should call investigationLoad successfully', function(done) {
        //uncomment below and update the code to test investigationLoad
        //instance.investigationLoad(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationLoadVersion', function() {
      it('should call investigationLoadVersion successfully', function(done) {
        //uncomment below and update the code to test investigationLoadVersion
        //instance.investigationLoadVersion(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('investigationUpdate', function() {
      it('should call investigationUpdate successfully', function(done) {
        //uncomment below and update the code to test investigationUpdate
        //instance.investigationUpdate(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
