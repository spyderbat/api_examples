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
    instance = new SpyderbatApi.Source();
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

  describe('Source', function() {
    it('should create an instance of Source', function() {
      // uncomment below and update the code to test Source
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be.a(SpyderbatApi.Source);
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property lastData (base name: "last_data")', function() {
      // uncomment below and update the code to test the property lastData
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property lastIngestChunkEndTime (base name: "last_ingest_chunk_end_time")', function() {
      // uncomment below and update the code to test the property lastIngestChunkEndTime
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property lastStoredChunkEndTime (base name: "last_stored_chunk_end_time")', function() {
      // uncomment below and update the code to test the property lastStoredChunkEndTime
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property orgUid (base name: "org_uid")', function() {
      // uncomment below and update the code to test the property orgUid
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property resourceName (base name: "resource_name")', function() {
      // uncomment below and update the code to test the property resourceName
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property resourcePolicy (base name: "resource_policy")', function() {
      // uncomment below and update the code to test the property resourcePolicy
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property runtimeDescription (base name: "runtime_description")', function() {
      // uncomment below and update the code to test the property runtimeDescription
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property runtimeDetails (base name: "runtime_details")', function() {
      // uncomment below and update the code to test the property runtimeDetails
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property uid (base name: "uid")', function() {
      // uncomment below and update the code to test the property uid
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property validFrom (base name: "valid_from")', function() {
      // uncomment below and update the code to test the property validFrom
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

    it('should have the property validTo (base name: "valid_to")', function() {
      // uncomment below and update the code to test the property validTo
      //var instance = new SpyderbatApi.Source();
      //expect(instance).to.be();
    });

  });

}));
