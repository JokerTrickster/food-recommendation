import { Suspense } from 'react';
import { RouterProvider } from 'react-router-dom';
import { router } from './router';
import FallbackComponent from '@shared/ui/fallback-component';

function App() {
  return (
    <Suspense fallback={<FallbackComponent />}>
      <RouterProvider router={router} />
    </Suspense>
  );
}

export default App;
