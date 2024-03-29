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
    instance = new SpyderbatApi.DashboardSearchCreateInput();
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

  describe('DashboardSearchCreateInput', function() {
    it('should create an instance of DashboardSearchCreateInput', function() {
      // uncomment below and update the code to test DashboardSearchCreateInput
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be.a(SpyderbatApi.DashboardSearchCreateInput);
    });

    it('should have the property  (base name: "")', function() {
      // uncomment below and update the code to test the property 
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property data (base name: "data")', function() {
      // uncomment below and update the code to test the property data
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property notify (base name: "notify")', function() {
      // uncomment below and update the code to test the property notify
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property notifyFrequency (base name: "notify_frequency")', function() {
      // uncomment below and update the code to test the property notifyFrequency
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property search (base name: "search")', function() {
      // uncomment below and update the code to test the property search
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new SpyderbatApi.DashboardSearchCreateInput();
      //expect(instance).to.be();
    });

  });

}));
