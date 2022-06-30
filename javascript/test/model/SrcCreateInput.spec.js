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
    factory(root.expect, root.Sbapi);
  }
}(this, function(expect, Sbapi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new Sbapi.SrcCreateInput();
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

  describe('SrcCreateInput', function() {
    it('should create an instance of SrcCreateInput', function() {
      // uncomment below and update the code to test SrcCreateInput
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be.a(Sbapi.SrcCreateInput);
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property runtimeDescription (base name: "runtime_description")', function() {
      // uncomment below and update the code to test the property runtimeDescription
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property runtimeDetails (base name: "runtime_details")', function() {
      // uncomment below and update the code to test the property runtimeDetails
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property uid (base name: "uid")', function() {
      // uncomment below and update the code to test the property uid
      //var instance = new Sbapi.SrcCreateInput();
      //expect(instance).to.be();
    });

  });

}));
