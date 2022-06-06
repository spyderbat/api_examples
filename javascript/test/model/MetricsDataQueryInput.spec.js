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
    instance = new Sbapi.MetricsDataQueryInput();
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

  describe('MetricsDataQueryInput', function() {
    it('should create an instance of MetricsDataQueryInput', function() {
      // uncomment below and update the code to test MetricsDataQueryInput
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be.a(Sbapi.MetricsDataQueryInput);
    });

    it('should have the property aggs (base name: "aggs")', function() {
      // uncomment below and update the code to test the property aggs
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property causal (base name: "causal")', function() {
      // uncomment below and update the code to test the property causal
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property contextUid (base name: "context_uid")', function() {
      // uncomment below and update the code to test the property contextUid
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property dataType (base name: "data_type")', function() {
      // uncomment below and update the code to test the property dataType
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property endTime (base name: "end_time")', function() {
      // uncomment below and update the code to test the property endTime
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property expr (base name: "expr")', function() {
      // uncomment below and update the code to test the property expr
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property ids (base name: "ids")', function() {
      // uncomment below and update the code to test the property ids
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property orgUid (base name: "org_uid")', function() {
      // uncomment below and update the code to test the property orgUid
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property projection (base name: "projection")', function() {
      // uncomment below and update the code to test the property projection
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property query (base name: "query")', function() {
      // uncomment below and update the code to test the property query
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property queryFrom (base name: "query_from")', function() {
      // uncomment below and update the code to test the property queryFrom
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property querySize (base name: "query_size")', function() {
      // uncomment below and update the code to test the property querySize
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property returnRptrs (base name: "return_rptrs")', function() {
      // uncomment below and update the code to test the property returnRptrs
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property rptrs (base name: "rptrs")', function() {
      // uncomment below and update the code to test the property rptrs
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property sortCol (base name: "sort_col")', function() {
      // uncomment below and update the code to test the property sortCol
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property sortOrder (base name: "sort_order")', function() {
      // uncomment below and update the code to test the property sortOrder
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property srcUid (base name: "src_uid")', function() {
      // uncomment below and update the code to test the property srcUid
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

    it('should have the property startTime (base name: "start_time")', function() {
      // uncomment below and update the code to test the property startTime
      //var instance = new Sbapi.MetricsDataQueryInput();
      //expect(instance).to.be();
    });

  });

}));