import Input from '@/components/ui/input';
import { StyleSheet, View, Text, ImageBackground } from 'react-native';
import CustomButton from '@/components/ui/button';
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withRepeat,
  withTiming,
  Easing,
} from 'react-native-reanimated';
import { useEffect } from 'react';

export default function LoginScreen() {
  const translateX = useSharedValue(0);
  const translateY = useSharedValue(0);

  useEffect(() => {
    translateX.value = withRepeat(
      withTiming(80, { duration: 8000, easing: Easing.inOut(Easing.sin) }),
      -1,
      true
    );
    translateY.value = withRepeat(
      withTiming(60, { duration: 6000, easing: Easing.inOut(Easing.sin) }),
      -1,
      true
    );
  }, []);

  const animatedStyle = useAnimatedStyle(() => ({
    transform: [
      { translateX: translateX.value - 300 },
      { translateY: translateY.value - 300 },
      { rotate: '-45deg' },
    ],
  }));

  return (
    <View style={styles.container}>
      <Animated.Image
        source={require('@/assets/images/doodle-bg.jpg')}
        resizeMode="repeat"
        style={[
          {
            position: 'absolute',
            width: '400%',
            height: '400%',
            opacity: 0.02,
            top: '-100%',
            left: 0,
          },
          animatedStyle,
        ]}
      />

      <View style={{
        flexDirection: 'column',
        justifyContent: 'center',
        height: '100%',
      }}>
        <Text style={[styles.title, { 
          marginBottom: 40, 
          textAlign: 'center',
        }]}>
          Welcome back to Popout!
        </Text>

        <View style={{ flexDirection: 'column', gap: 4 }}>
          <Input
            styleView={styles.input}
            labelColor="black"
            placeholder="hello@example.com"
          />
          <Input
            styleView={styles.input}
            labelColor="black"
            placeholder="Enter your password"
          />
        </View>

        <View style={{ marginTop: 120, gap: 10 }}>
          <CustomButton
            style={styles.button}
            onPress={() => {}}
            title="Continue"
            iconsName="arrow-right"
          />
          <CustomButton
            style={styles.buttonSecondary}
            onPress={() => {}}
            title="Continue with Google"
            colorTitle="black"
          />
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 30,
    paddingTop: 50,
    backgroundColor: 'white',
  },
  title: {
    fontSize: 50,
    fontWeight: 'bold',
    color: 'black',
  },
  input: {
    width: '100%',
    marginTop: 20,
  },
  button: {
    backgroundColor: '#8422DC',
  },
  buttonSecondary: {
    backgroundColor: '#E2E2E2',
  },
});