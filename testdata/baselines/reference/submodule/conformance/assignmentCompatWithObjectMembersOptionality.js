//// [tests/cases/conformance/types/typeRelationships/assignmentCompatibility/assignmentCompatWithObjectMembersOptionality.ts] ////

//// [assignmentCompatWithObjectMembersOptionality.ts]
// Derived member is not optional but base member is, should be ok

class Base { foo: string; }
class Derived extends Base { bar: string; }
class Derived2 extends Derived { baz: string; }

module TargetHasOptional {
    // targets
    interface C {
        opt?: Base
    }
    var c: C;

    var a: { opt?: Base; }
    var b: typeof a = { opt: new Base() }

    // sources
    interface D {
        opt: Base;
    }
    interface E {
        opt: Derived;
    }
    interface F {
        opt?: Derived;
    }
    var d: D;
    var e: E;
    var f: F;

    // all ok
    c = d;
    c = e;
    c = f;
    c = a;

    a = d;
    a = e;
    a = f;
    a = c;

    b = d;
    b = e;
    b = f;
    b = a;
    b = c;
}

module SourceHasOptional {
    // targets
    interface C {
        opt: Base
    }
    var c: C;

    var a: { opt: Base; }
    var b = { opt: new Base() }

    // sources
    interface D {
        opt?: Base;
    }
    interface E {
        opt?: Derived;
    }
    interface F {
        opt: Derived;
    }
    var d: D;
    var e: E;
    var f: F;

    c = d; // error
    c = e; // error
    c = f; // ok
    c = a; // ok

    a = d; // error
    a = e; // error
    a = f; // ok
    a = c; // ok

    b = d; // error
    b = e; // error
    b = f; // ok
    b = a; // ok
    b = c; // ok
}

//// [assignmentCompatWithObjectMembersOptionality.js]
// Derived member is not optional but base member is, should be ok
class Base {
    foo;
}
class Derived extends Base {
    bar;
}
class Derived2 extends Derived {
    baz;
}
var TargetHasOptional;
(function (TargetHasOptional) {
    var c;
    var a;
    var b = { opt: new Base() };
    var d;
    var e;
    var f;
    // all ok
    c = d;
    c = e;
    c = f;
    c = a;
    a = d;
    a = e;
    a = f;
    a = c;
    b = d;
    b = e;
    b = f;
    b = a;
    b = c;
})(TargetHasOptional || (TargetHasOptional = {}));
var SourceHasOptional;
(function (SourceHasOptional) {
    var c;
    var a;
    var b = { opt: new Base() };
    var d;
    var e;
    var f;
    c = d; // error
    c = e; // error
    c = f; // ok
    c = a; // ok
    a = d; // error
    a = e; // error
    a = f; // ok
    a = c; // ok
    b = d; // error
    b = e; // error
    b = f; // ok
    b = a; // ok
    b = c; // ok
})(SourceHasOptional || (SourceHasOptional = {}));
