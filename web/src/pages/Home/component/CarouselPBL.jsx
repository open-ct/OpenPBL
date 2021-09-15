import React from 'react';
import BannerAnim from 'rc-banner-anim';
import {Image} from 'antd';
import TweenOne, {TweenOneGroup} from 'rc-tween-one';
import 'rc-banner-anim/assets/index.css';

import './carousel-pbl.less';

const {Element, Arrow, Thumb} = BannerAnim;
const BgElement = Element.BgElement;

class CarouselPBL extends React.Component {
  constructor(props) {
    super(props);
    this.imgArray = [
      {
        img: 'https://cdn.open-ct.com/task-resources/lbaf23/carousel-top-1.jpg',
        title: 'O p e n P B L',
        text: 'Project Based Learning',
        titleStyle: {color: 'white', fontSize: '60px', fontWeight: 'bold', fontFamily: 'Times New Roman'},
        textStyle: {color: 'white', fontSize: '30px', marginTop: '50px'}
      },
      {
        img: 'https://cdn.open-ct.com/task-resources/lbaf23/carousel-top-2.jpg',
        title: 'O p e n P B L',
        text: 'Propose - Plan - Execute - Judge',
        titleStyle: {color: 'lightblue', fontSize: '60px', fontWeight: 'bold', fontFamily: 'Times New Roman'},
        textStyle: {color: 'lightblue', fontSize: '30px', marginTop: '50px'}
      },
      {
        img: 'https://cdn.open-ct.com/task-resources/lbaf23/carousel-top-3.jpg',
        title: 'O p e n P B L',
        text: 'explore real-world problems',
        titleStyle: {color: 'black', fontSize: '60px', fontWeight: 'bold', fontFamily: 'Times New Roman'},
        textStyle: {color: 'black', fontSize: '30px', marginTop: '50px'}
      },
    ];
    this.state = {
      intShow: 0,
      prevEnter: false,
      nextEnter: false,
      thumbEnter: false,
    };
    [
      'onChange',
      'prevEnter',
      'prevLeave',
      'nextEnter',
      'nextLeave',
      'onMouseEnter',
      'onMouseLeave',
    ].forEach((method) => this[method] = this[method].bind(this));
  }

  onChange(type, int) {
    if (type === 'before') {
      this.setState({
        intShow: int,
      });
    }
  }

  getNextPrevNumber() {
    let nextInt = this.state.intShow + 1;
    let prevInt = this.state.intShow - 1;
    if (nextInt >= this.imgArray.length) {
      nextInt = 0;
    }
    if (prevInt < 0) {
      prevInt = this.imgArray.length - 1;
    }

    return [prevInt, nextInt];
  }

  prevEnter() {
    this.setState({
      prevEnter: true,
    });
  }

  prevLeave() {
    this.setState({
      prevEnter: false,
    });
  }

  nextEnter() {
    this.setState({
      nextEnter: true,
    });
  }

  nextLeave() {
    this.setState({
      nextEnter: false,
    });
  }

  onMouseEnter() {
    this.setState({
      thumbEnter: true,
    });
  }

  onMouseLeave() {
    this.setState({
      thumbEnter: false,
    });
  }

  render() {
    const intArray = this.getNextPrevNumber();
    const thumbChildren = this.imgArray.map((item, i) =>
      (
        <span key={i.toString()}>
          <Image src={item.img}/>
        </span>
      ));
    return (
      <BannerAnim
        onChange={this.onChange}
        onMouseEnter={this.onMouseEnter}
        onMouseLeave={this.onMouseLeave}
        prefixCls="custom-arrow-thumb"
        autoPlay
      >
        {this.imgArray.map((item, index) => (
          <Element key={index.toString()} prefixCls="banner-user-elem">
            <BgElement
              key="bg"
              className="bg"
              style={{
                backgroundImage: `url(${item.img})`,

                backgroundSize: 'cover',
                backgroundPosition: 'center',
              }}
            />
            <TweenOne className="banner-user-title" animation={{y: 30, opacity: 0, type: 'from'}}>
              <div style={item.titleStyle}>{item.title}</div>
            </TweenOne>
            <TweenOne className="banner-user-text" animation={{y: 30, opacity: 0, type: 'from', delay: 100}}>
              <div style={item.textStyle}>{item.text}</div>
            </TweenOne>
          </Element>
        ))}

        <Arrow
          arrowType="prev"
          key="prev"
          prefixCls="user-arrow"
          component={TweenOne}
          onMouseEnter={this.prevEnter}
          onMouseLeave={this.prevLeave}
          animation={{left: this.state.prevEnter ? 0 : -120}}
        >
          <div className="arrow"/>
          <TweenOneGroup
            enter={{opacity: 0, type: 'from'}}
            leave={{opacity: 0}}
            appear={false}
            className="img-wrapper"
            component="ul"
          >
            <li style={{backgroundImage: `url(${this.imgArray[intArray[0]].img})`}} key={intArray[0]}/>
          </TweenOneGroup>
        </Arrow>
        <Arrow
          arrowType="next"
          key="next"
          prefixCls="user-arrow"
          component={TweenOne}
          onMouseEnter={this.nextEnter}
          onMouseLeave={this.nextLeave}
          animation={{right: this.state.nextEnter ? 0 : -120}}
        >
          <div className="arrow"/>
          <TweenOneGroup
            enter={{opacity: 0, type: 'from'}}
            leave={{opacity: 0}}
            appear={false}
            className="img-wrapper"
            component="ul"
          >
            <li style={{backgroundImage: `url(${this.imgArray[intArray[1]].img})`}} key={intArray[1]}/>
          </TweenOneGroup>
        </Arrow>

        <Thumb
          prefixCls="user-thumb"
          key="thumb"
          component={TweenOne}
          animation={{bottom: this.state.thumbEnter ? 0 : -70}}
        >
          {thumbChildren}
        </Thumb>
      </BannerAnim>
    );
  }
}

export default CarouselPBL;
