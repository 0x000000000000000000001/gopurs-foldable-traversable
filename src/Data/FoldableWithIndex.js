export const foldlWithIndexImpl = function (f) {
  return function (z) {
    return function (xs) {
      var acc = z;
      var len = xs.length;
      for (var i = 0; i < len; i++) {
        acc = f(i)(acc)(xs[i]);
      }
      return acc;
    };
  };
};

export const foldrWithIndexImpl = function (f) {
  return function (z) {
    return function (xs) {
      var acc = z;
      var len = xs.length;
      for (var i = len - 1; i >= 0; i--) {
        acc = f(i)(xs[i])(acc);
      }
      return acc;
    };
  };
};
