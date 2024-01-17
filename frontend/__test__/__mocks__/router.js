// mocks/router.js
import { useRouter as useRouterOriginal } from 'next/router';

// const useRouter = () => ({
//   ...useRouterOriginal(),
//   pathname: '/your-test-path',
//   query: { id: '123' }, 
//   push: jest.fn(), 
//   events: {
//     on: jest.fn(),
//     off: jest.fn()
//   },
//   beforePopState: jest.fn(() => null),
//   prefetch: jest.fn(() => null)
// });

const mockRouter = () => ({
    useRouter() {
      return {
        route: '/seats',
        pathname: '',
        query: '',
        asPath: '',
        push: jest.fn(),
        events: {
          on: jest.fn(),
          off: jest.fn()
        },
        beforePopState: jest.fn(() => null),
        prefetch: jest.fn(() => null)
      };
    },
  })
export { useRouter,mockRouter };