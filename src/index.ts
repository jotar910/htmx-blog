import '../scss/main.scss';

// Components
import './components/pinterest-layout';
import './components/netflix-horizontal-list';
import './components/carousel';

// TW elements
import {
  Carousel,
  Input,
  initTE,
} from "tw-elements";

initTE({ Carousel, Input });

console.debug('running application...');
