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
    instance = new SpyderbatApi.NotificationPolicyDestination();
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

  describe('NotificationPolicyDestination', function() {
    it('should create an instance of NotificationPolicyDestination', function() {
      // uncomment below and update the code to test NotificationPolicyDestination
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be.a(SpyderbatApi.NotificationPolicyDestination);
    });

    it('should have the property data (base name: "data")', function() {
      // uncomment below and update the code to test the property data
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property email (base name: "email")', function() {
      // uncomment below and update the code to test the property email
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property orgUid (base name: "org_uid")', function() {
      // uncomment below and update the code to test the property orgUid
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property slack (base name: "slack")', function() {
      // uncomment below and update the code to test the property slack
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property users (base name: "users")', function() {
      // uncomment below and update the code to test the property users
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

    it('should have the property webhook (base name: "webhook")', function() {
      // uncomment below and update the code to test the property webhook
      //var instance = new SpyderbatApi.NotificationPolicyDestination();
      //expect(instance).to.be();
    });

  });

}));
