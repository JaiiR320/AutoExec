import clsx from 'clsx';
import Heading from '@theme/Heading';
import styles from './styles.module.css';

const FeatureList = [
  {
    title: 'Easy to Use',
    Svg: require('@site/static/img/console_icon.svg').default,
    description: (
      <>
        An Autoexec file can either be created in the terminal or
        through a text editor, and added to the bucket directory later!
      </>
    ),
  },
  {
    title: 'Quickly Launch your Favorite Programs',
    Svg: require('@site/static/img/computer_icon.svg').default,
    description: (
      <>
        Autoexec allows you to launch your favorite programs on startup
        and create multiple configurations for different situations.
      </>
    ),
  },
  {
    title: 'Quick to Edit',
    Svg: require('@site/static/img/settings_icon.svg').default,
    description: (
      <>
        Edit your Autoexec file with ease, and add new programs to your list
        with a single command!
      </>
    ),
  },
];

function Feature({Svg, title, description}) {
  return (
    <div className={clsx('col col--4')}>
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures() {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
