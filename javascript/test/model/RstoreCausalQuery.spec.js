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
    instance = new Sbapi.RstoreCausalQuery();
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

  describe('RstoreCausalQuery', function() {
    it('should create an instance of RstoreCausalQuery', function() {
      // uncomment below and update the code to test RstoreCausalQuery
      //var instance = new Sbapi.RstoreCausalQuery();
      //expect(instance).to.be.a(Sbapi.RstoreCausalQuery);
    });

    it('should have the property peer (base name: "peer")', function() {
      // uncomment below and update the code to test the property peer
      //var instance = new Sbapi.RstoreCausalQuery();
      //expect(instance).to.be();
    });

    it('should have the property post (base name: "post")', function() {
      // uncomment below and update the code to test the property post
      //var instance = new Sbapi.RstoreCausalQuery();
      //expect(instance).to.be();
    });

    it('should have the property pre (base name: "pre")', function() {
      // uncomment below and update the code to test the property pre
      //var instance = new Sbapi.RstoreCausalQuery();
      //expect(instance).to.be();
    });

  });

}));
