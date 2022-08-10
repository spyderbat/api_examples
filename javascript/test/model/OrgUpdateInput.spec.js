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
    instance = new SpyderbatApi.OrgUpdateInput();
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

  describe('OrgUpdateInput', function() {
    it('should create an instance of OrgUpdateInput', function() {
      // uncomment below and update the code to test OrgUpdateInput
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be.a(SpyderbatApi.OrgUpdateInput);
    });

    it('should have the property activeSources (base name: "active_sources")', function() {
      // uncomment below and update the code to test the property activeSources
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property activeUsers (base name: "active_users")', function() {
      // uncomment below and update the code to test the property activeUsers
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property orgTypeUid (base name: "org_type_uid")', function() {
      // uncomment below and update the code to test the property orgTypeUid
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property ownerEmail (base name: "owner_email")', function() {
      // uncomment below and update the code to test the property ownerEmail
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property ownerUid (base name: "owner_uid")', function() {
      // uncomment below and update the code to test the property ownerUid
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property resourceName (base name: "resource_name")', function() {
      // uncomment below and update the code to test the property resourceName
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property resourcePolicy (base name: "resource_policy")', function() {
      // uncomment below and update the code to test the property resourcePolicy
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property totalSources (base name: "total_sources")', function() {
      // uncomment below and update the code to test the property totalSources
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property totalUsers (base name: "total_users")', function() {
      // uncomment below and update the code to test the property totalUsers
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property validFrom (base name: "valid_from")', function() {
      // uncomment below and update the code to test the property validFrom
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

    it('should have the property validTo (base name: "valid_to")', function() {
      // uncomment below and update the code to test the property validTo
      //var instance = new SpyderbatApi.OrgUpdateInput();
      //expect(instance).to.be();
    });

  });

}));
