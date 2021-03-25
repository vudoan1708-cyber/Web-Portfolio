import { provide, inject } from '@vue/runtime-core';

export function provideStore(storeName, store) {
  provide(storeName, store);
}

export function useStore(storeName) {
  const store = inject(storeName);
  let err = null;
  if (!store) {
    err = 'No Store Provided';
  }

  return { store, err };
}
