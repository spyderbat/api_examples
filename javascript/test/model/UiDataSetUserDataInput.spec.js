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
    instance = new Sbapi.UiDataSetUserDataInput();
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

  describe('UiDataSetUserDataInput', function() {
    it('should create an instance of UiDataSetUserDataInput', function() {
      // uncomment below and update the code to test UiDataSetUserDataInput
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be.a(Sbapi.UiDataSetUserDataInput);
    });

    it('should have the property  (base name: "")', function() {
      // uncomment below and update the code to test the property 
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property data (base name: "data")', function() {
      // uncomment below and update the code to test the property data
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property lastUsed (base name: "last_used")', function() {
      // uncomment below and update the code to test the property lastUsed
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property tags (base name: "tags")', function() {
      // uncomment below and update the code to test the property tags
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property uid (base name: "uid")', function() {
      // uncomment below and update the code to test the property uid
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property validFrom (base name: "valid_from")', function() {
      // uncomment below and update the code to test the property validFrom
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

    it('should have the property validTo (base name: "valid_to")', function() {
      // uncomment below and update the code to test the property validTo
      //var instance = new Sbapi.UiDataSetUserDataInput();
      //expect(instance).to.be();
    });

  });

}));
