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
    instance = new SpyderbatApi.Expr();
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

  describe('Expr', function() {
    it('should create an instance of Expr', function() {
      // uncomment below and update the code to test Expr
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be.a(SpyderbatApi.Expr);
    });

    it('should have the property and (base name: "and")', function() {
      // uncomment below and update the code to test the property and
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property containsStr (base name: "contains_str")', function() {
      // uncomment below and update the code to test the property containsStr
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property equals (base name: "equals")', function() {
      // uncomment below and update the code to test the property equals
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property exists (base name: "exists")', function() {
      // uncomment below and update the code to test the property exists
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property greaterThan (base name: "greater_than")', function() {
      // uncomment below and update the code to test the property greaterThan
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property hasPrefix (base name: "has_prefix")', function() {
      // uncomment below and update the code to test the property hasPrefix
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property hasSuffix (base name: "has_suffix")', function() {
      // uncomment below and update the code to test the property hasSuffix
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property _in (base name: "in")', function() {
      // uncomment below and update the code to test the property _in
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property lessThan (base name: "less_than")', function() {
      // uncomment below and update the code to test the property lessThan
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property not (base name: "not")', function() {
      // uncomment below and update the code to test the property not
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property or (base name: "or")', function() {
      // uncomment below and update the code to test the property or
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property property (base name: "property")', function() {
      // uncomment below and update the code to test the property property
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property reMatch (base name: "re_match")', function() {
      // uncomment below and update the code to test the property reMatch
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property schema (base name: "schema")', function() {
      // uncomment below and update the code to test the property schema
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

    it('should have the property timeRange (base name: "time_range")', function() {
      // uncomment below and update the code to test the property timeRange
      //var instance = new SpyderbatApi.Expr();
      //expect(instance).to.be();
    });

  });

}));
