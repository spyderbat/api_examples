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
    instance = new Sbapi.ApiKeyCreateInput();
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

  describe('ApiKeyCreateInput', function() {
    it('should create an instance of ApiKeyCreateInput', function() {
      // uncomment below and update the code to test ApiKeyCreateInput
      //var instance = new Sbapi.ApiKeyCreateInput();
      //expect(instance).to.be.a(Sbapi.ApiKeyCreateInput);
    });

    it('should have the property description (base name: "description")', function() {
      // uncomment below and update the code to test the property description
      //var instance = new Sbapi.ApiKeyCreateInput();
      //expect(instance).to.be();
    });

    it('should have the property validTo (base name: "valid_to")', function() {
      // uncomment below and update the code to test the property validTo
      //var instance = new Sbapi.ApiKeyCreateInput();
      //expect(instance).to.be();
    });

  });

}));
