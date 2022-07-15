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
    instance = new SpyderbatApi.UIData();
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

  describe('UIData', function() {
    it('should create an instance of UIData', function() {
      // uncomment below and update the code to test UIData
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be.a(SpyderbatApi.UIData);
    });

    it('should have the property data (base name: "data")', function() {
      // uncomment below and update the code to test the property data
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property dataKey (base name: "data_key")', function() {
      // uncomment below and update the code to test the property dataKey
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property lastUsed (base name: "last_used")', function() {
      // uncomment below and update the code to test the property lastUsed
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property orgUid (base name: "org_uid")', function() {
      // uncomment below and update the code to test the property orgUid
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property ownerUid (base name: "owner_uid")', function() {
      // uncomment below and update the code to test the property ownerUid
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property sourceUid (base name: "source_uid")', function() {
      // uncomment below and update the code to test the property sourceUid
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property uid (base name: "uid")', function() {
      // uncomment below and update the code to test the property uid
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property validFrom (base name: "valid_from")', function() {
      // uncomment below and update the code to test the property validFrom
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

    it('should have the property validTo (base name: "valid_to")', function() {
      // uncomment below and update the code to test the property validTo
      //var instance = new SpyderbatApi.UIData();
      //expect(instance).to.be();
    });

  });

}));
